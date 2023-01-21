.PHONY:list

platform-db:
	@docker-compose -f docker-compose.yml --profile db up

platform-down:
	@docker-compose -f docker-compose.yml down 

list:
	@echo "available commands"
	
terra-init:
	@terraform init

terra-apply:
	@terraform apply	

terra-destroy:
	@terraform destroy	