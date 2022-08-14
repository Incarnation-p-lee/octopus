GO                      :=$(if $(V), go, @go)
RM                      :=$(if $(V), rm -rf, @rm -rf)
CAT                     :=$(if $(V), cat, @cat)
CD                      :=$(if $(V), cd, @cd)
ECHO                    :=$(if $(V), echo, @echo)

package_prefix          :=github.com/Incarnation-p-lee/octopus/
coverage_file           :=coverage.txt

go_tmp                  :=$(shell find . -name "*.go" | grep -v octopus.go)
go_dirs                 :=$(sort $(dir $(go_tmp)))     # will remove duplicated.
go_coverages_files      :=$(addsuffix $(coverage_file), $(go_dirs))
go_ldflags              :="-s"

cmd_dir                 :=cmd/octopus
cmd_main                :=$(cmd_dir)/octopus.go
cmd                     :=$(subst .go, , $(cmd_main))

TARGET                  :=$(cmd_dir)/octopus

.PHONY: cmd test clean release

$(cmd):$(cmd_main)
	$(ECHO) "Build     $(TARGET)"
	$(CD) $(cmd_dir) && go build $(notdir $(cmd_main)) && cd -> /dev/null

test:$(go_coverages_files)
	$(ECHO) "Test      completed"

release:$(cmd_main)
	$(ECHO) "Release   $(TARGET)"
	$(CD) $(cmd_dir) && go build -ldflags $(go_ldflags) $(notdir $(cmd_main)) && cd -> /dev/null

$(go_coverages_files):%/coverage.txt:%
	$(ECHO) "Test      $<"
	$(CD) $< && $(run_test_cases) && cd -> /dev/null
	$(CAT) $@ >> coverage.txt
	$(RM) $@

clean:
	$(RM) $(cmd) $(go_coverages_files) $(coverage_file)

define run_test_cases
    go test -covermode=atomic -coverprofile=$(notdir $@) \
        $(if $(filter cmd/octopus, $<), octopus_test.go octopus.go, $(get_test_package_path))
endef

define get_test_package_path
    $(if $(filter pkg/%, $<), $(package_prefix), )$<
endef
