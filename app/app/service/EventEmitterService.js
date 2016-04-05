System.register(["angular2/core"], function(exports_1, context_1) {
    "use strict";
    var __moduleName = context_1 && context_1.id;
    var core_1;
    var EmitterService;
    return {
        setters:[
            function (core_1_1) {
                core_1 = core_1_1;
            }],
        execute: function() {
            EmitterService = (function () {
                function EmitterService() {
                }
                EmitterService.get = function (channel) {
                    if (!this._emitters[channel])
                        this._emitters[channel] = new core_1.EventEmitter();
                    return this._emitters[channel];
                };
                EmitterService._emitters = {};
                return EmitterService;
            }());
            exports_1("EmitterService", EmitterService);
        }
    }
});
//# sourceMappingURL=EventEmitterService.js.map