export default {

    // function converts fields from the input/edit objects to the correct data types
    convert(object, fields) {
        // make a copy of the object
        let modify = Object.assign({}, object);

        // find which fields require conversions
        for (let key in modify) {
            for (let i = 0; i < fields.length; i++) {
                // match key with list of fields to find type property
                if (fields[i].key == key) {
                    
                    // convert string to number
                    if (fields[i].type == "number") {
                        modify[key] = Number(modify[key]);
                    }
                    // convert dates to iso string
                    else if (fields[i].type == "date") {
                        let date = new Date(modify[key]);

                        modify[key] = date.toISOString();
                    }   
                }
            }
        }
        return modify;
    },

    emptyFields(input) {
        let inputs = Object.values(input);

        for (let value of inputs) {
            if (value == "") {
                return true;
            }
        }
        return false;
    }
}