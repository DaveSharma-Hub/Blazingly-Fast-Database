const convertToType = (value, type) => {
    switch(type){
        case "string":
            return String(value);
        case "integer":
            return Number.parseInt(value);
        case "float":
            return Number.parseFloat(value);
        default:
            return String(value);
    }
}


const unmarshall = (marshalledJson) => {
    if("value" in marshalledJson && "type" in marshalledJson && typeof(marshalledJson)==="object"){
        return convertToType(marshalledJson["value"],marshalledJson["type"]);
    }else{
        const newObj = {};
        Object.entries(marshalledJson).forEach(([key,value])=>{
            newObj[key] = unmarshall(value);
        });
        return newObj;
    }
}



module.exports = {
    unmarshall:unmarshall
}