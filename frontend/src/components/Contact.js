import {
  FaInstagram,
  FaFacebook,
  FaTwitter,
  FaTumblr,
  FaPinterest,
} from "react-icons/fa";

const Contact = () => {
  return (
    <footer className="bg-gray-500 bg-opacity-30 text-black py-10 px-5">
      <h3 className="font-bold text-2xl mb-2">Contact Us</h3>
      <div className="container mx-auto grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 text-sm">
        {/* Shop Section */}
        <div className="flex flex-col items-center md:items-start border-t border-gray-700 pt-5">
          <h3 className="font-bold mb-2">Shop</h3>
          <ul className="space-y-2 text-center md:text-left">
            <li>Gift Guides</li>
            <li>Fan Art</li>
            <li>New Works</li>
            <li>Blog</li>
            <li>Student Discount</li>
            <li>Login</li>
            <li>Signup</li>
            <li>Bulk Orders</li>
          </ul>
        </div>

        {/* About Section */}
        <div className="flex flex-col items-center md:items-start border-t border-gray-700 pt-5">
          <h3 className="font-bold mb-2">About</h3>
          <ul className="space-y-2 text-center md:text-left">
            <li>About Us</li>
            <li>Social Responsibility</li>
            <li>Investor Center</li>
            <li>Partner Program</li>
            <li>Affiliates</li>
            <li>Sell your art</li>
            <li>Jobs</li>
            <li>Artist Blog</li>
          </ul>
        </div>

        {/* Help Section */}
        <div className="flex flex-col items-center md:items-start border-t border-gray-700 pt-5">
          <h3 className="font-bold mb-2">Help</h3>
          <ul className="space-y-2 text-center md:text-left">
            <li>Delivery</li>
            <li>Returns</li>
            <li>Help Center</li>
            <li>Guidelines</li>
            <li>Product Safety</li>
            <li>Copyright</li>
            <li>Contact Us</li>
            <li>Cookie Settings</li>
          </ul>
        </div>

        {/* Social Section */}
        <div className="flex flex-col items-center md:items-start border-t border-gray-700 pt-5">
          <h3 className="font-bold mb-2">Social</h3>
          <ul className="space-y-2 flex flex-col items-center">
            <li className="flex items-center gap-2">
              <FaInstagram /> Instagram
            </li>
            <li className="flex items-center gap-2">
              <FaFacebook /> Facebook
            </li>
            <li className="flex items-center gap-2">
              <FaTwitter /> Twitter
            </li>
            <li className="flex items-center gap-2">
              <FaTumblr /> Tumblr
            </li>
            <li className="flex items-center gap-2">
              <FaPinterest /> Pinterest
            </li>
          </ul>
        </div>
      </div>

      {/* Bottom Section */}
      <div className="mt-10 border-t border-gray-700 pt-5 text-center text-xs">
        <p>© RoastWear. All Rights Reserved</p>
        <div className="flex justify-center gap-4 mt-3 flex-wrap">
          <p>User Agreement</p>
          <p>Privacy Policy</p>
          <p>Do not sell my personal information</p>
          <p>Cookie Policy</p>
        </div>
      </div>
    </footer>
  );
};

export default Contact;
