import React from 'react';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from './Form.css';

class Form extends React.Component {
  constructor(props) {
    super(props);
    this.onOK = this.onOK.bind(this);
    this.onCancel = this.onCancel.bind(this);
  }

  onOK(event) {
    event.preventDefault();
    if (this.props.onOKAction) {
      this.props.dispatch(this.props.onOKAction(this.props.formData));
    }
  }

  onCancel(event) {
    event.preventDefault();
    if (this.props.onCancelAction) {
      this.props.dispatch(this.props.onCancelAction);
    }
  }

  render() {
    return (
      <div styleName="FormContainer">
        <div styleName="Form">
          <div styleName="FormTitleDiv">
            <p styleName="FormTitleText">{this.props.title}</p>
          </div>
          <hr styleName="FormHr" />
          <div styleName="FormContentDiv">
            {this.props.children}
          </div>
          <hr styleName="FormHr" />
          <div styleName="FormButtonDiv">
            <button onClick={this.onCancel} styleName="FormButtonText">Cancel</button >
            <button onClick={this.onOK} styleName="FormButtonText">OK</button >
          </div>
        </div>
      </div>
    );
  }
}

Form.propTypes = {
  dispatch: PropTypes.func,
  title: PropTypes.string,
  children: PropTypes.oneOfType([PropTypes.object, PropTypes.array]),
  onCancelAction: PropTypes.func,
  onOKAction: PropTypes.func,
  onCancel: PropTypes.func,
  onOK: PropTypes.func,
  formData: PropTypes.object,
};

export default connect(null, null)(CSSModules(Form, styles));
