import {
  FaInstagram,
  FaFacebook,
  FaTwitter,
  FaTumblr,
  FaPinterest,
} from "react-icons/fa";

const Contact = () => {
  return (
    <footer className="bg-[#0b0c0f] text-gray-300 py-20 px-5 border-t border-gray-700">
      <h3 className="font-bold text-2xl mb-2 text-center md:text-left text-white">Contact Us</h3>
      <div className="container mx-auto grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 text-sm">
        {/* Shop Section */}
        <div className="flex flex-col items-center md:items-start pt-5">
          <h3 className="font-bold mb-2 text-[#25aae1]">Shop</h3>
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
        <div className="flex flex-col items-center md:items-start pt-5">
          <h3 className="font-bold mb-2 text-[#25aae1]">About</h3>
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
        <div className="flex flex-col items-center md:items-start pt-5">
          <h3 className="font-bold mb-2 text-[#25aae1]">Help</h3>
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
        <div className="flex flex-col items-center md:items-start pt-5">
          <h3 className="font-bold mb-2 text-[#25aae1]">Social</h3>
          <ul className="space-y-2 flex flex-col items-center md:items-start">
            <li className="flex items-center gap-2">
              <FaInstagram className="text-[#E1306C]" /> Instagram
            </li>
            <li className="flex items-center gap-2">
              <FaFacebook className="text-[#3b5998]" /> Facebook
            </li>
            <li className="flex items-center gap-2">
              <FaTwitter className="text-[#1DA1F2]" /> Twitter
            </li>
            <li className="flex items-center gap-2">
              <FaTumblr className="text-[#35465C]" /> Tumblr
            </li>
            <li className="flex items-center gap-2">
              <FaPinterest className="text-[#E60023]" /> Pinterest
            </li>
          </ul>
        </div>
      </div>

      {/* Bottom Section */}
      <div className="mt-10 pt-5 text-center text-xs text-gray-400 border-t border-gray-700">
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
