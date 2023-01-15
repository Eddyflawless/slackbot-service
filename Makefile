.PHONY:list

list:
	@echo "available commands"
	
terra-init:
	@terraform init


terra-apply:
	@terraform apply	

terra-destroy:
	@terraform destroy	