import React from 'react';
import PropTypes from 'prop-types';
import Dialog from '../../components/common/Dialog';

class AddHardwareDialog extends React.Component {
  constructor(props) {
    super(props);
    this.onOk = this.onOk.bind(this);
    this.onInputChange = this.onInputChange.bind(this);
    this.onCancel = this.onCancel.bind(this);
    this.state = {
      input: null
    };
  }

  onInputChange(event) {
    event.preventDefault();
    this.setState({
      input: event.target.value
    });
  }

  onOk(event) {
    event.preventDefault();
    this.props.onOk(this.state.input);
  }

  onCancel(event) {
    event.preventDefault();
    this.props.onCancel();
  }

  render() {
    const getChildren = () => {
      return (
        <input type="text" onChange={this.onInputChange} />
      );
    };

    return (
      <Dialog title="Add Hardware" children={getChildren()} onOk={this.onOk} onCancel={this.onCancel} />
    );
  }
}

AddHardwareDialog.propTypes = {
  onCancel: PropTypes.func,
  onOk: PropTypes.func
};

export default AddHardwareDialog;
