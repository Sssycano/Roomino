import axios from "axios";

const API_URL = "http://localhost:3000/";

const register = (username, password, first_name, last_name, email, phone, dob, gender) => {
  return axios.post(API_URL + "register", {
    username,
    password,
    first_name,
    last_name,
    email,
    phone,
    dob,
    gender: parseInt(gender),
    
  });
};


const login = (userName, password) => {
  return axios
    .post(API_URL + "login", {
      userName,
      password,
    })
    .then((response) => {
      if (response.data.token) {
        localStorage.setItem("user", JSON.stringify(response.data));
      }

      return response.data;
    });
};

const logout = () => {
  localStorage.removeItem("user");
};

const getCurrentUser = () => {
  return JSON.parse(localStorage.getItem("user"));
};

const AuthService = {
  register,
  login,
  logout,
  getCurrentUser,
};

export default AuthService;
