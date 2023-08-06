import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import styles from './(*)';

class {}{%%} extends Component{

  constructor(props) {
    super(props)
    this.state = {  }
  }

  render() {
    return(
        <div className={styles.container}>
          <div>{}{%%}</div>
          <div>{this.props.example}</div>
        </div>
    )
  }
}

{}{%%}.propTypes = {
  example: PropTypes.string.isRequired
};

function mapStateToProps(state) {
  return {
    example: state.example.text
  }
}

export default connect(mapStateToProps)({}{%%});
