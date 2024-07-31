import React from "react";
import {
  FaFacebookF,
  FaTwitter,
  FaInstagram,
  FaLinkedinIn,
} from "react-icons/fa";

const Footer = () => {
  return (
    <footer className="bg-gray-900 text-white py-10">
      <div className="container mx-auto flex flex-col md:flex-row justify-between items-center">
        {/* Brand Information and Social Links */}
        <div className="flex flex-col items-center md:items-start mb-6 md:mb-0">
          <h2 className="text-3xl font-bold mb-4">Car Sales</h2>
          <p className="text-center md:text-left max-w-md mb-4">
            We are a dedicated team committed to providing top-notch car sales
            services. Our mission is to offer reliable and affordable vehicles
            to meet your needs.
          </p>
          <div className="flex space-x-4">
            <a
              href="https://facebook.com"
              className="text-white hover:text-blue-500 transition-colors"
            >
              <FaFacebookF size={24} />
            </a>
            <a
              href="https://twitter.com"
              className="text-white hover:text-blue-400 transition-colors"
            >
              <FaTwitter size={24} />
            </a>
            <a
              href="https://instagram.com"
              className="text-white hover:text-pink-500 transition-colors"
            >
              <FaInstagram size={24} />
            </a>
            <a
              href="https://linkedin.com"
              className="text-white hover:text-blue-700 transition-colors"
            >
              <FaLinkedinIn size={24} />
            </a>
          </div>
        </div>

        {/* Important Links */}
        <div className="flex flex-col items-center md:items-end">
          <h2 className="text-2xl font-bold mb-4">Important Links</h2>
          <ul className="text-center md:text-right space-y-2">
            <li>
              <a href="/#" className="hover:underline">
                Home
              </a>
            </li>
            <li>
              <a href="/#cars" className="hover:underline">
                Cars
              </a>
            </li>
            <li>
              <a href="/#about" className="hover:underline">
                About
              </a>
            </li>
            <li>
              <a href="/#contact" className="hover:underline">
                Contact
              </a>
            </li>
          </ul>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
