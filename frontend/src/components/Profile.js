import React from "react";
import { Navigate } from 'react-router-dom';
import AuthService from "../services/auth.service";
import "./Profile.css"; 

const Profile = () => {
  const currentUser = AuthService.getCurrentUser();

  const handleSearch = () => {
    window.location.href = '/profile/unitinfo';
  };

  const handleUpdatePet = () => {
    window.location.href = '/profile/updatepet';
  };
  const handleRegisterPet = () => {
    window.location.href = '/profile/registerpet';
  };

  if(!currentUser) {
    return <Navigate to="/login" replace={true} />
  }

  return (
    <div className="container">
      <header className="jumbotron">
        <h3>
          <strong>{currentUser.username}</strong> Profile
        </h3>
      </header>
      <p>
        <strong>Token:</strong> {currentUser.token ? `${currentUser.token.substring(0, 20)} ... ${currentUser.token.substr(currentUser.token.length - 20)}` : 'No token available'}
      </p>
      <p>
        <strong>userName: </strong> {currentUser.user_name ? currentUser.user_name : 'No username available'}
      </p>
      <p>
        <strong>Email: </strong> {currentUser.email ? currentUser.email : 'No email available'}
      </p>
      <button className="btn btn-primary custom-btn" onClick={handleSearch}>Search Certain Apartment Units</button><br /><br />
      <button className="btn btn-primary custom-btn" onClick={handleRegisterPet}>Register Pet</button><br /><br />
      <button className="btn btn-primary custom-btn" onClick={handleUpdatePet}>Update Pet</button>
    </div>
  );
};

export default Profile;
