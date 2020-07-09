const utils = {
    getOnInputOrOnChangeProp(fn) {
        return this.isIE11() ? { onInput: fn } : { onChange: fn };
    },

    isIE11() {
        return !!window.MSInputMethodContext && !!document.documentMode;
    },
};

export default utils;