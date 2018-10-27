import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import { connect } from 'react-redux';
// import CSSModules from 'react-css-modules';
import ExpansionPanel from '@material-ui/core/ExpansionPanel';
import ExpansionPanelDetails from '@material-ui/core/ExpansionPanelDetails';
import ExpansionPanelSummary from '@material-ui/core/ExpansionPanelSummary';
import Typography from '@material-ui/core/Typography';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
// import styles from './App.css';

const styles = theme => ({
  root: {
    margin: '2em'
  },
  heading: {
    fontSize: theme.typography.pxToRem(15),
    flexBasis: '33.33%',
    flexShrink: 0,
  },
  detail: {
    display: 'flex',
    flexDirection: 'row',
    flexWrap: 'wrap',
    justifyContent: 'flex-start',
    alignItems: 'flex-start',
    alignContent: 'flex-start'
  },
  columnLeft: {
    flexBasis: '33.33%',
  },
  columnRight: {
    flexBasis: '66.66%',
  },
  secondaryHeading: {
    fontSize: theme.typography.pxToRem(15),
    color: theme.palette.text.secondary,
  }
});

class EnclosureDetail extends React.Component {
  state = {
    expanded: 'basic'
  };

  handleChange = panel => (event, expanded) => {
    this.setState({
      expanded: expanded ? panel : false,
    });
  };

  render() {
    const { classes, enclosure } = this.props;
    const { expanded } = this.state;

    return (
      <div className={classes.root}>
        <ExpansionPanel expanded={expanded === 'basic'} onChange={this.handleChange('basic')}>
          <ExpansionPanelSummary expandIcon={<ExpandMoreIcon/>}>
            <Typography className={classes.heading}>Basic</Typography>
          </ExpansionPanelSummary>
          <ExpansionPanelDetails className={classes.detail}>
            <div className={classes.columnLeft}>
              <Typography className={classes.secondaryHeading}>Name</Typography>
            </div>
            <div className={classes.columnRight}>
              <Typography>{enclosure.Name}</Typography>
            </div>

            <div className={classes.columnLeft}>
              <Typography className={classes.secondaryHeading}>Description</Typography>
            </div>
            <div className={classes.columnRight}>
              <Typography>{enclosure.Description}</Typography>
            </div>

            <div className={classes.columnLeft}>
              <Typography className={classes.secondaryHeading}>Name</Typography>
            </div>
            <div className={classes.columnRight}>
              <Typography>{enclosure.Name}</Typography>
            </div>

            <div className={classes.columnLeft}>
              <Typography className={classes.secondaryHeading}>State</Typography>
            </div>
            <div className={classes.columnRight}>
              <Typography>{enclosure.State + '/' + enclosure.StateReason}</Typography>
            </div>

            <div className={classes.columnLeft}>
              <Typography className={classes.secondaryHeading}>Created At</Typography>
            </div>
            <div className={classes.columnRight}>
              <Typography>{enclosure.CreatedAt}</Typography>
            </div>

            <div className={classes.columnLeft}>
              <Typography className={classes.secondaryHeading}>Updated At</Typography>
            </div>
            <div className={classes.columnRight}>
              <Typography>{enclosure.UpdatedAt}</Typography>
            </div>
          </ExpansionPanelDetails>
        </ExpansionPanel>
      </div>
    );
  }
}

EnclosureDetail.propTypes = {
  classes: PropTypes.object.isRequired,
  enclosure: PropTypes.object,
};

// export default connect()(CSSModules(EnclosureDetail, styles, {allowMultiple: true}));
export default connect()(withStyles(styles)(EnclosureDetail));
