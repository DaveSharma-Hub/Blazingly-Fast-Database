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
};


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
};

const arrayToObjectConversion = (array) => {
    return array.reduce((acc,curr)=>{
        acc[curr[0]] = {
            value:convertToType(curr[1], curr[2]),
            type:curr[2]
        } 
        return acc;
    },{});
};

const objectToArrayConversion = (obj) => {
    return Object.entries(obj).map(([key,value])=>{
        return [key, String(value.value), value.type];
    });
}

module.exports = {
    unmarshall:unmarshall,
    arrayToObjectConversion:arrayToObjectConversion,
    objectToArrayConversion: objectToArrayConversion
}