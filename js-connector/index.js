const axios = require('axios');
const { unmarshall } = require('./util/utilFunctions');

class BFDB {
    comparators = ["EQUAL"];

    constructor(databaseEndpoint){
        this.databaseEndpoint = databaseEndpoint;
    }
    async createTable(tableName, headerArray){
        try{
            const endpoint = `${this.databaseEndpoint}/createTable`;
            await axios.post(endpoint,{
                tableName:tableName,
                headerArray:headerArray
            });
        }catch(e){
            console.log(e);
        }
    }

    async addDataToTable(tableName, partitionKey, payload){
        try{
            const endpoint = `${this.databaseEndpoint}/addData`;
            const result = await axios.post(endpoint,{
                table_name:tableName,
                partition_key:partitionKey,
                payload: payload
            });
            console.log(result.data);
        }catch(e){
            console.log(e);
        }
    }
    async updateDataToTable(tableName, partitionKey, payload){
        try{
            const endpoint = `${this.databaseEndpoint}/updateData`;
            const result = await axios.post(endpoint,{
                table_name:tableName,
                partition_key:partitionKey,
                payload: payload
            });
            console.log(result.data);
        }catch(e){
            console.log(e);
        }
    }

    async getDataFromTable(tableName,partitionKey){
        try{
            const endpoint = `${this.databaseEndpoint}/queryData`;
            const result = await axios.post(endpoint,{
                table_name:tableName,
                partition_key:partitionKey
            });
            return unmarshall(JSON.parse(result.data));
        }catch(e){
            console.log(e);
        }
    }

    async removeDataFromTable(tableName,key){
        try{
            const endpoint = `${this.databaseEndpoint}/removeData`;
            await axios.post(endpoint,{
                table_name:tableName,
                partition_key:key
            });
        }catch(e){
            console.log(e);
        }
    }

    async scanDataFromTable(tableName, innerKey, innerValue, comparator){
        if(!this.comparators.includes(comparator)){
            throw Error(`Comparator ${comparator} doesnt exist in ${this.comparators}`);
        }
        try{
            const endpoint = `${this.databaseEndpoint}/scanData`;
            const result = await axios.post(endpoint,{
                table_name:tableName,
                inner_key:innerKey,
                inner_value:innerValue,
                comparator:comparator
            });
            return unmarshall(JSON.parse(result.data));
        }catch(e){
            console.log(e);
        }
    }

    async test(){
        try{
            const endpoint = `${this.databaseEndpoint}/test`
            const result = await axios.get(endpoint);
            return unmarshall(JSON.parse(result.data));
        }catch(e){
            console.log(e);
        }   
    }

}

module.exports ={
    BFDB:BFDB
}