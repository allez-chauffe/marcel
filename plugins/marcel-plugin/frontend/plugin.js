'use strict';

var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }

function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }

var Text = function (_Marcel$Plugin) {
  _inherits(Text, _Marcel$Plugin);

  function Text() {
    _classCallCheck(this, Text);

    var _this = _possibleConstructorReturn(this, (Text.__proto__ || Object.getPrototypeOf(Text)).call(this, {
      defaultProps: {
        text: '',
        stylesvar: {}
      }
    }));

    _this.content = document.getElementById('content');
    return _this;
  }

  _createClass(Text, [{
    key: 'render',
    value: function render() {
      var _props = this.props,
          text = _props.text,
          stylesvar = _props.stylesvar;


      this.content.innerText = text;

      if (stylesvar['primary-color']) this.p.style.color = stylesvar['primary-color'];
      if (stylesvar['font-family']) this.p.style.fontFamily = stylesvar['font-family'];
    }
  }]);

  return Text;
}(Marcel.Plugin);

var instance = new Text();