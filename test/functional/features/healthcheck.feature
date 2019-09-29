Feature: Healthcheck

  Scenario: Application is running
	  When I send "GET" request to "/_health"
	  Then The response code should be 200
