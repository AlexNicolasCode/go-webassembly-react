class InitWasm {
    constructor (
        initInstance
    ) {}

    async getInstance () {
        if (!this.initInstance) {
            const go = new Go()
            const { instance } = await WebAssembly.instantiateStreaming(
                fetch("http://localhost:8000/main.wasm"),
                go.importObject,
            )
            this.initInstance = instance
            return this
        }
        return this.initInstance
    }
    
    getMethods () {
        const add = this.initInstance.exports.add;
        return { add }
    }
}

export const initWasm = await new InitWasm().getInstance()