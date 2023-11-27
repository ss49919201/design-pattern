interface Validtor {
    valid(value: String): Boolean;
}

class CustomRegExp {
    value: String
    constructor(value: String) {
        this.value = value;
    };
    match():Boolean {
        const regExp = /^S/
        return regExp.test(this.value.toString());
    };
}

class ValidatorAdapter {
    valid(value: String): Boolean {
        const customRegExp = new CustomRegExp(value);
        return customRegExp.match();
    };
}

const validator: Validtor = new ValidatorAdapter();
console.log(validator.valid('SSS'));
console.log(validator.valid('ASS'));
