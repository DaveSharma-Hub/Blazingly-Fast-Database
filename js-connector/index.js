const axios = require('axios');
const { unmarshall } = require('./util/utilFunctions');

class BFDB {
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

    async addDataToTable(tableName, data){
        try{
            const endpoint = `${this.databaseEndpoint}/addData`;
            await axios.post(endpoint,{
                tableName:tableName,
                tableData:data
            });
        }catch(e){
            console.log(e);
        }
    }

    async getDataFromTable(tableName,dataId){
        try{
            const endpoint = `${this.databaseEndpoint}/queryData`;
            await axios.post(endpoint,{
                tableName:tableName,
                dataId:dataId
            });
        }catch(e){
            console.log(e);
        }
    }
    async removeDataFromTable(tableName,dataId){
        try{
            const endpoint = `${this.databaseEndpoint}/removeData`;
            await axios.post(endpoint,{
                tableName:tableName,
                dataId:dataId
            });
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