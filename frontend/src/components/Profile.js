import React from "react";
import { Navigate } from 'react-router-dom';
import AuthService from "../services/auth.service";

const Profile = () => {
  const currentUser = AuthService.getCurrentUser();

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
    </div>
  );
};

export default Profile;
