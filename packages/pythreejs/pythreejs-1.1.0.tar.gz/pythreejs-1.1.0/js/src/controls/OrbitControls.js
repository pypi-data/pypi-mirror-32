var _ = require('underscore');
var OrbitControls = require('../examples/controls/OrbitControls.js').OrbitControls;
var OrbitControlsAutogen = require('./OrbitControls.autogen');


var OrbitControlsModel = OrbitControlsAutogen.OrbitControlsModel.extend({

    constructThreeObject: function() {
        var controlling = this.get('controlling');
        var obj = new OrbitControls(controlling.obj);
        obj.dispose();  // Disconnect events, we need to (dis-)connect on freeze/thaw
        obj.enableKeys = false; // turn off keyboard navigation

        return obj;
    },

    setupListeners: function() {
        OrbitControlsAutogen.OrbitControlsModel.prototype.setupListeners.call(this);
        var that = this;
        this.obj.addEventListener('end', function() {
            that.update_controlled();
        });
    },

    update_controlled: function() {
        // Since OrbitControls changes the position of the object, we
        // update the position when we've stopped moving the object
        // it's probably prohibitive to update it in real-time
        var controlling = this.get('controlling');
        var pos = controlling.obj.position;
        var qat = controlling.obj.quaternion;
        controlling.set(
            {
                position: [pos.x, pos.y, pos.z],
                quaternion: [qat._x, qat._y, qat._z, qat._w],
                zoom: controlling.obj.zoom,
            },
            'pushFromThree'
        );
        controlling.save_changes();
    },

});

module.exports = {
    OrbitControlsModel: OrbitControlsModel,
};
