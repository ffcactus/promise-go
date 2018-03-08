import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from '../../styles/ServerFrame.css';


function ObjectEthernetInterface(props) {
  const IPv4Addresses = props.ethernetInterface.IPv4Addresses;
  const IPv6Addresses = props.ethernetInterface.IPv6Addresses;
  if (IPv4Addresses.length === 0) {
    IPv4Addresses.push({
      'Address': 'None',
      'SubnetMask': 'None',
      'Gateway': 'None'
    });
  }
  if (IPv6Addresses.length === 0) {
    IPv6Addresses.push({
      'Address': 'None',
      'PrefixLength': 'None',
      'AddressState': 'None'
    });
  }
  const ipv4RowSpan = (IPv4Addresses.length + 1).toString();
  const ipv6RowSpan = (IPv6Addresses.length + 1).toString();
  return (<table>
    <tbody>
      <tr>
        <th styleName="level1">Name</th>
        <td colSpan="3">{props.ethernetInterface.Name}</td>
      </tr>
      <tr>
        <th styleName="level1">MAC</th>
        <td colSpan="3">{props.ethernetInterface.PermanentMACAddress}</td>
      </tr>
      <tr>
        <th styleName="level1" rowSpan={ipv4RowSpan}>IPv4</th>
        <th styleName="level2">Address</th>
        <th styleName="level2">Subnet</th>
        <th styleName="level2">Gateway</th>
      </tr>
      {IPv4Addresses.map((each, i) => {
        return (
          <tr key={i.toString()}>
            <td>{each.Address}</td>
            <td>{each.SubnetMask}</td>
            <td>{each.Gateway}</td>
          </tr>
        );
      })}
      <tr>
        <th styleName="level1" rowSpan={ipv6RowSpan}>IPv6</th>
        <th styleName="level2">Address</th>
        <th styleName="level2">Prefix length</th>
        <th styleName="level2">State</th>
      </tr>
      {IPv6Addresses.map((each, i) => {
        return (
          <tr key={i.toString()}>
            <td>{each.Address}</td>
            <td>{each.PrefixLength}</td>
            <td>{each.AddressState}</td>
          </tr>
        );
      })}
      <tr>
        <th>VLan</th>
        <td colSpan="3">{JSON.stringify(props.ethernetInterface.VLANs)}</td>
      </tr>
    </tbody>
  </table>);
}

ObjectEthernetInterface.propTypes = {
  ethernetInterface: PropTypes.object,
};

export default CSSModules(ObjectEthernetInterface, styles);
