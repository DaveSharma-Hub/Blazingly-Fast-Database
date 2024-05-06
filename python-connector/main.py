# class BFDB:
#     comparators = ["EQUAL"]

#     def __init__(self,databaseEndpoint):
#         self.databaseEndpoint = databaseEndpoint
    
#     def createTable(self,tableName, headerArray):
#         try:
#             endpoint = f'{self.databaseEndpoint}/createTable'
#             axios.post(endpoint,{
#                 "tableName":tableName,
#                 "headerArray":headerArray
#             })
#         except(e):
#             print(e)

    
#     def addDataToTable(self,tableName, partitionKey, payload):
#         try:
#             endpoint = f'{self.databaseEndpoint}/addData'
#             result = axios.post(endpoint,{
#                 "table_name":tableName,
#                 "partition_key":partitionKey,
#                 "payload": payload
#             })
#         except(e):
#             print(e)

#     def updateDataToTable(self, tableName, partitionKey, payload):
#         try:
#             endpoint = f'{self.databaseEndpoint}/updateData'
#             result = axios.post(endpoint,{
#                 "table_name":tableName,
#                 "partition_key":partitionKey,
#                 "payload": payload
#             })
#         except(e):
#             print(e)

#     def getDataFromTable(self, tableName,partitionKey):
#         try:
#             endpoint = f'{self.databaseEndpoint}/queryData'
#             result = axios.post(endpoint,{
#                 "table_name":tableName,
#                 "partition_key":partitionKey
#             })
#             return unmarshall(JSON.parse(result.data));
#         except(e):
#             print(e)

#     def removeDataFromTable(self, tableName,key):
#         try:
#             endpoint = f'{self.databaseEndpoint}/removeData'
#             axios.post(endpoint,{
#                 "table_name":tableName,
#                 "partition_key":key
#             })
#         except(e):
#             print(e)

#     def scanDataFromTable(self, tableName, innerKey, innerValue, comparator):
#         if(comparator not in self.comparators):
#             raise Warning(f'Comparator {comparator} doesnt exist in {self.comparators}')
        
#         try:
#             endpoint = f'{self.databaseEndpoint}/scanData'
#             result =  axios.post(endpoint,{
#                 "table_name":tableName,
#                 "inner_key":innerKey,
#                 "inner_value":innerValue,
#                 "comparator":comparator
#             })
#             return unmarshall(JSON.parse(result.data));
#         except(e):
#             print(e)

