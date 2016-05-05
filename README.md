# sofe2016a

Hello, and welcome to the place where magical things happen!

Today we'll be working on an API written in the Go programming language. I hope you're all set up and ready to GO!

"Code without testing is like a taco without meat". You will have a set of manual test cases to try out on your API, that all need to pass and give the proper results.

"Don't leave broken windows" There's a famous experiment where a car was left abandoned. This car was never vandalized until one single window was broken. After that, it was completely destroyed.

If you see a bug, FIX IT!! If you see a code smell, FIX IT!!

HINT: Remember the DRY principle: Don't Repeat Yourself

Here are the cards on your story board.

[BUG] Oh no! the users of can't sign in! Please fix the sign in or the world will end! Okay maybe it won't end but we need to be able to get new users.

[Feature] The users of our API wish to be able to add and track their balance. This number should be part of their profile when they sign up. If the users lack the balance, an error should be shown, and they should not be allowed to sign up. When there is a GET request for a user with the userID, it should show the users balance.

[FEATURE] Please make a new endpoint that will add or substract a users' balance. Ensure that the response has the users' profile and updated balance.

Here are the system requirements!!

Requirements:

*user should not be given an access token with an incorrect password to log-in

*user should be able to login with username and password.

*user should not be able to sign up without a username or password.

*user should not be able to sign up without a balance.

*user should be able to update their username, and log in with their updated username.

*user should be able to add to their balance or substract to their balance.

BONUS: Add unit tests to the methods.
