const getValue = (name: string) => {
    switch (name) {
        case "yaml":
            return `name: zero\nage: 18`
        case "json":
            return `{
    "code": 200,
    "data": [
        {
            "CreatedAt": "2023-01-30T09:56:58.543+08:00",
            "DeletedAt": null,
            "ID": 4,
            "UpdatedAt": "2023-01-30T09:56:58.543+08:00",
            "animal_name": "test3",
            "animal_type": "普通动物",
            "image": "60e74100a04171edbddc472180b80101",
            "introduce": "123123123",
            "latin_name": "test",
            "province": "河北",
            "video_id": "28df6ac0b8cc4683b66e8440f16e5998"
        }
    ],
    "msg": "查询成功"
}`
        case "toml":
            let tomlValue = `name="zero"\nage=18`
            return tomlValue
        case "hcl":
            let hclValue = `data {\n\tname = "zero"\n\tage = 18\n}`
            return hclValue
        case "env":
            let envValue = `database=localhost\nport=3306\n`
            return envValue
        case "ini":
            let iniValue = `[database]\nname = zero\nage = 18`
            return iniValue
        default:
            return "未知语言，请提交issue"
    }
}
export {
    getValue
}