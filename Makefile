DOCKER_SERVICE_NAME=masterdata
DOCKER_REPOSITORY=projects/smartsustainability/masterdata

include $(abspath ../../../../common/makefiles/go/Makefile)


.PHONY: proto
proto: ## Build the proto of the contracts
	@printf "Building proto...\n"
	@${DC_WORKSPACE} "\
		protoc -I ./ -I ./../../../hyperion/kore/ --kore2_out=. --go_out=. --go_opt=paths=source_relative --kore2_opt=paths=source_relative **/**/*.proto && \
        			./contract/scripts/handle-proto-versions.sh && \
		protoc -I ./ -I ./../../../hyperion/kore/ --kore2_out=. --go_out=. --go_opt=paths=source_relative --kore2_opt=paths=source_relative **/**/*.proto && \
			./contract/scripts/handle-proto-versions.sh && \
		protoc -I ./ -I ./../../../hyperion/kore/ --kresolver2_out=. --go_out=. --go_opt=paths=source_relative --kresolver2_opt=paths=source_relative **/**/*.proto \
	"
