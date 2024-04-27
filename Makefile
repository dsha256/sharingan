.DEFAULT_GOAL = help

temporal: ## Run Temporal client.
	temporal server start-dev --dynamic-config-value frontend.enableUpdateWorkflowExecution=true

help: ## Prints this message.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	sort | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: help temporal
