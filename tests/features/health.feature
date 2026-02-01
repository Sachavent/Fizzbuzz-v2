Feature: Health endpoint
  Scenario: Check health status
    When I make a get request on the route "/health"
    Then the response status code is 200
    And the response body is:
      """
      {
        "status": "ok"
      }
      """
