interface Validator {
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

// CustomRegExp は Validator を満たさないので、橋渡しになる Class を定義する
class ValidatorAdapter {
    valid(value: String): Boolean {
        const customRegExp = new CustomRegExp(value);
        return customRegExp.match();
    };
}

const validator: Validator = new ValidatorAdapter();
console.log(validator.valid('SSS'));
console.log(validator.valid('ASS'));
