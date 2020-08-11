# Folder Forex


	Lamda function named 'forex' gets data from DynamoDB and Api is deployed which returns the data obtained.
	Api is deployed link https://lktkd58l71.execute-api.ap-south-1.amazonaws.com/staging/forex.
	CORS is handled.
	Helper functions for Client Side and Server Side are written.
	Lambda function to get a particular conversion Forex item is written.



# Folder SaveForex:


	Lambda function saveForex gets data from 3rd Party Forex Api (Link :   https://www.freeforexapi.com/Home/Api )
	and saves the data to dynamodb using putItem function.
	The saveforex lambda function runs every min using the AWS EventBridge and updates the data in dynamodb.  
