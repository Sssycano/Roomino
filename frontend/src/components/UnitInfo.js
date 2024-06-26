import React, { useState } from 'react';
import Form from 'react-validation/build/form';
import Input from 'react-validation/build/input';
import Button from 'react-validation/build/button';
import { Link } from 'react-router-dom'; 
import UserService from '../services/user.service';

const UnitInfo = () => {
  const [formData, setFormData] = useState({
    companyName: '',
    buildingName: ''
  });
  const [unitInfo, setUnitInfo] = useState([]);
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState('');

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    setMessage('');
    setLoading(true);

    UserService.searchUnits(formData.companyName, formData.buildingName)
      .then(data => {
        setUnitInfo(data.data);
        setLoading(false);
      })
      .catch(error => {
        setMessage('Error retrieving unit info');
        setLoading(false);
      });
  };

  return (
    <div className="container">
      <div className="row">
        <div className="col-md-6 offset-md-3">
          <h2>Search Certain Apartment Units</h2>
          <Form onSubmit={handleSubmit}>
            <div className="form-group">
              <label htmlFor="companyName">Company Name</label>
              <Input
                type="text"
                className="form-control"
                name="companyName"
                value={formData.companyName}
                onChange={handleChange}
              />
            </div>
            <div className="form-group">
              <label htmlFor="buildingName">Building Name</label>
              <Input
                type="text"
                className="form-control"
                name="buildingName"
                value={formData.buildingName}
                onChange={handleChange}
              />
            </div>
            <div className="form-group">
              <Button className="btn btn-primary" disabled={loading}>Search</Button>
            </div>
          </Form>
          {message && <div className="alert alert-danger">{message}</div>}
          <hr />
          <div>
            {unitInfo.length > 0 && (
              <div>
                <h3>Unit Information</h3>
                <ul>
                  {unitInfo.map(unit => (
                    <li key={unit.unit_rent_id}>
                      <strong>Unit Rent ID:</strong> 
                      <Link to={`/profile/interests?unit_rent_id=${unit.unit_rent_id}`}>{unit.unit_rent_id}</Link> 
                      <br />
                      <strong>Monthly Rent:</strong> {unit.monthly_rent}<br />
                      <strong>Square Footage:</strong> {unit.square_footage}<br />
                      <strong>Available Date for Move-in:</strong> {new Date(unit.available_date_for_move_in).toLocaleDateString('en-US', {
                          year: 'numeric',
                          month: '2-digit',
                          day: '2-digit'
                        })}<br />
                      <strong>Your pet allowed? :</strong> {unit.is_pet_allowed ? 'Yes' : 'No'}
                      <hr />
                    </li>
                  ))}
                </ul>
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};

export default UnitInfo;
