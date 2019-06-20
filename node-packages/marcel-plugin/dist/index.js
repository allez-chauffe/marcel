'use strict';

var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };

var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

;(function () {
  var state = void 0,
      props = void 0;

  var Plugin = function () {
    _createClass(Plugin, [{
      key: 'propsDidChange',
      value: function propsDidChange() {}
    }, {
      key: 'render',
      value: function render() {}
    }]);

    function Plugin() {
      var _this = this;

      var defaults = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : {};

      _classCallCheck(this, Plugin);

      state = defaults.defaultState;
      props = defaults.defaultProps;

      this.setState = function (newState) {
        state = _extends({}, state, newState);
        setTimeout(function () {
          return _this.render();
        });
      };

      addEventListener('message', function (event) {
        if (event.source !== parent) return;
        var message = event.data;

        if (message.type === 'propsChange') {
          var _message$payload = message.payload,
              newProps = _message$payload.newProps,
              prevProps = _message$payload.prevProps;

          props = newProps;
          setTimeout(function () {
            _this.render();
            setTimeout(function () {
              return _this.propsDidChange(prevProps || {});
            });
          });
        }
      });

      parent.postMessage({ type: 'loaded' }, '*');
    }

    _createClass(Plugin, [{
      key: 'props',
      get: function get() {
        return props;
      }
    }, {
      key: 'state',
      get: function get() {
        return state;
      }
    }]);

    return Plugin;
  }();

  var Debug = function () {
    function Debug() {
      _classCallCheck(this, Debug);
    }

    _createClass(Debug, null, [{
      key: 'changeProps',
      value: function changeProps(newProps, prevProps) {
        dispatchEvent(new MessageEvent('message', {
          source: parent,
          data: {
            type: 'propsChange',
            payload: { newProps: newProps, prevProps: prevProps }
          }
        }));
      }
    }]);

    return Debug;
  }();

  window.Marcel = { Plugin: Plugin, Debug: Debug };
})();