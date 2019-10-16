export function removePrefix(data, tableName) {
    let json = data;
    let prefix = tableName + '.';

    for (let key in json) {
        if (key.startsWith(prefix)) {
            let replacedKey = key.replace(prefix, '');

            json[replacedKey] = json[key];
            delete json[key];
        }
    }

    return json;
}

export function addPrefix(data, tableName) {
    let json = data;
    let prefix = tableName + '.';

    for (let key in json) {
        let replacedKey = prefix + key;

        json[replacedKey] = json[key];
        delete json[key];
    }

    return json;
}