Feature: Fizzbuzz endpoint
    Scenario: Should return 400 when missing int1 parameter
        When I make a get request on the route "/fizzbuzz/result?int2=5&limit=10&str1=fizz&str2=buzz"
        Then the response status code is 400
        And the response body is:
            """
            {
                "errors": [
                    "Key: 'GetResultQuery.Int1' Error:Field validation for 'Int1' failed on the 'required' tag"
                ]
            }
            """

    Scenario: Should return 400 when int1 is under 0
        When I make a get request on the route "/fizzbuzz/result?int1=-1&int2=5&limit=10&str1=fizz&str2=buzz"
        Then the response status code is 400
        And the response body is:
            """
            {
                "errors": [
                    "Key: 'GetResultQuery.Int1' Error:Field validation for 'Int1' failed on the 'gt' tag"
                ]
            }
            """

    Scenario: Should return 400 when missing int2 parameter
        When I make a get request on the route "/fizzbuzz/result?int1=5&limit=10&str1=fizz&str2=buzz"
        Then the response status code is 400
        And the response body is:
            """
            {
                "errors": [
                    "Key: 'GetResultQuery.Int2' Error:Field validation for 'Int2' failed on the 'required' tag"
                ]
            }
            """

    Scenario: Should return 400 when int2 is under 0
        When I make a get request on the route "/fizzbuzz/result?int1=5&int2=-1&limit=10&str1=fizz&str2=buzz"
        Then the response status code is 400
        And the response body is:
            """
            {
                "errors": [
                    "Key: 'GetResultQuery.Int2' Error:Field validation for 'Int2' failed on the 'gt' tag"
                ]
            }
            """

    Scenario: Should return 400 when missing limit parameter
        When I make a get request on the route "/fizzbuzz/result?int1=5&int2=5&str1=fizz&str2=buzz"
        Then the response status code is 400
        And the response body is:
            """
            {
                "errors": [
                    "Key: 'GetResultQuery.Limit' Error:Field validation for 'Limit' failed on the 'required' tag"
                ]
            }
            """

    Scenario: Should return 400 when limit is greater than 10000
        When I make a get request on the route "/fizzbuzz/result?int1=5&int2=5&limit=10001&str1=fizz&str2=buzz"
        Then the response status code is 400
        And the response body is:
            """
            {
                "errors": [
                    "Key: 'GetResultQuery.Limit' Error:Field validation for 'Limit' failed on the 'lte' tag"
                ]
            }
            """

    Scenario: Should return 400 when missing str1 parameter
        When I make a get request on the route "/fizzbuzz/result?int1=5&int2=5&limit=10&str2=buzz"
        Then the response status code is 400
        And the response body is:
            """
            {
                "errors": [
                    "Key: 'GetResultQuery.Str1' Error:Field validation for 'Str1' failed on the 'required' tag"
                ]
            }
            """

    Scenario: Should return 400 when str1 is longer than 10 characters
        When I make a get request on the route "/fizzbuzz/result?int1=5&int2=5&limit=10&str1=fizzbuzzfizzbuzz&str2=buzz"
        Then the response status code is 400
        And the response body is:
            """
            {
                "errors": [
                    "Key: 'GetResultQuery.Str1' Error:Field validation for 'Str1' failed on the 'max' tag"
                ]
            }
            """

    Scenario: Should return 400 when missing str2 parameter
        When I make a get request on the route "/fizzbuzz/result?int1=5&int2=5&limit=10&str1=fizz"
        Then the response status code is 400
        And the response body is:
            """
            {
                "errors": [
                    "Key: 'GetResultQuery.Str2' Error:Field validation for 'Str2' failed on the 'required' tag"
                ]
            }
            """

    Scenario: Should return 400 when str2 is longer than 10 characters
        When I make a get request on the route "/fizzbuzz/result?int1=5&int2=5&limit=10&str1=fizz&str2=fizzbuzzfizzbuzz"
        Then the response status code is 400
        And the response body is:
            """
            {
                "errors": [
                    "Key: 'GetResultQuery.Str2' Error:Field validation for 'Str2' failed on the 'max' tag"
                ]
            }
            """

    Scenario: Should return 400 when missing int1 and int2 parameter
        When I make a get request on the route "/fizzbuzz/result?limit=10&str1=fizz&str2=buzz"
        Then the response status code is 400
        And the response body is:
            """
            {
                "errors": [
                    "Key: 'GetResultQuery.Int1' Error:Field validation for 'Int1' failed on the 'required' tag",
                    "Key: 'GetResultQuery.Int2' Error:Field validation for 'Int2' failed on the 'required' tag"
                ]
            }
            """

    Scenario: Should return 400 when int1 is not an integer
        When I make a get request on the route "/fizzbuzz/result?int1=abc&int2=5&limit=10&str1=fizz&str2=buzz"
        Then the response status code is 400
        And the response body is:
            """
            {
                "error": "invalid query parameters"
            }
            """

    Scenario: Check fizzbuzz result
        When I make a get request on the route "/fizzbuzz/result?int1=2&int2=5&limit=10&str1=fizz&str2=buzz"
        Then the response status code is 200
        And the response body is:
            """
            1,fizz,3,fizz,buzz,fizz,7,fizz,9,fizzbuzz
            """
