@todos
Feature: Todos

	Scenario: Create Todo | Validation errors
		Given I set header "Content-Type" with value "application/json"
		When I send "POST" request to "/todos" with body:
		"""
		{}
		"""
		Then The response code should be 422
		And The response should be a valid json
		And The response should match json:
			"""
			{
			   "code":"VALIDATION_FAILED",
			   "message":"Validation Failed",
			   "fields":[
				  {
					 "code":"REQUIRED",
					 "message":"Description is a required field",
					 "location":"description"
				  }
			   ]
			}
			"""

	Scenario: Create Todo | Success
		Given I set header "Content-Type" with value "application/json"
		When I send "POST" request to "/todos" with body:
		"""
		{
			"description": "test todo"
		}
		"""
		Then The response code should be 201
		And The response should be a valid json
		And The response should match json schema "todo.json"

	Scenario: List Todos
		Given I have todos:
			| Todo 1 |
			| Todo 2 |
		When I send "GET" request to "/todos"
		Then The response code should be 200
		And  The response should match json schema "todos_list.json"
		And The json path "$" should have count "2"
