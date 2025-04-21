import React from "react";
import { motion } from "framer-motion";
import { Swiper, SwiperSlide } from "swiper/react";
import "swiper/css";
import "swiper/css/effect-fade";
import { EffectFade, Autoplay } from "swiper/modules";

const features = [
  {
    title: "👕 Premium Quality",
    desc: "Top-grade fabrics and long-lasting custom prints made with care.",
  },
  {
    title: "🚀 Lightning Delivery",
    desc: "Shipped within 48 hours, straight to your door. No delays.",
  },
  {
    title: "🌱 Eco Friendly",
    desc: "We print on-demand using sustainable, water-based inks.",
  },
];

const About = () => {
  return (
    <div className="bg-[#0b0c0f] text-gray-300 px-6 py-12 w-full min-h-screen space-y-20">

      {/* Top Quote */}
      <motion.div
        initial={{ opacity: 0, y: 30 }}
        whileInView={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.5 }}
        viewport={{ once: true }}
        className="text-center max-w-3xl mx-auto text-xl font-semibold"
      >
        <p>“Wearing a funny T-shirt is cheaper than therapy… and way more stylish.”</p>
      </motion.div>

      {/* Features Section */}
      <motion.div
        initial={{ opacity: 0, scale: 0.9 }}
        whileInView={{ opacity: 1, scale: 1 }}
        transition={{ duration: 0.7, delay: 0.2 }}
        viewport={{ once: true }}
        className="grid md:grid-cols-3 gap-6"
      >
        {features.map((item, idx) => (
          <div
            key={idx}
            className="bg-[#121417] shadow-lg rounded-xl p-6 border border-[#1f2937] hover:shadow-xl transition duration-300"
          >
            <h3 className="text-xl font-semibold mb-2 text-white">{item.title}</h3>
            <p className="text-[#9ca3af]">{item.desc}</p>
          </div>
        ))}
      </motion.div>

      {/* Tharun Bhascker Quote Section */}
      <motion.div
        initial={{ opacity: 0, y: 30 }}
        whileInView={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.6, delay: 0.1 }}
        viewport={{ once: true }}
        className="flex flex-col items-center text-center space-y-4"
      >
        <img
          src="/Assets/tharun.jpg"
          alt="Tharun Bhascker"
          className="w-full max-w-md h-auto object-cover rounded-xl shadow-[0_0_50px_#ffd369]"
        />

        <blockquote className="bg-[#121417]/90 text-[#25aae1] px-4 py-3 rounded-xl shadow-[0_0_10px_#25aae155] max-w-xl text-lg font-semibold relative backdrop-blur">
          “Rey Karthik... vaallaki size L cheppu! Custom print lu kooda cheppu ra veskundam!”
          <span className="block mt-2 text-sm text-[#9ca3af] italic">
            – Not really Tharun Bhascker… but sounds like him.
          </span>
        </blockquote>
      </motion.div>

      {/* Image Slider */}
  {/* Frontend Section */}
<motion.div
  initial={{ opacity: 0 }}
  whileInView={{ opacity: 1 }}
  transition={{ duration: 0.8 }}
  viewport={{ once: true }}
  className="max-w-4xl mx-auto"
>
  <h2 className="text-2xl font-bold mb-2 text-center text-white">Frontend Team</h2>
  <Swiper
    spaceBetween={30}
    effect={"fade"}
    modules={[EffectFade, Autoplay]}
    autoplay={{ delay: 3000 }}
    loop={true}
    className="rounded-lg overflow-hidden shadow-lg mb-10"
  >
    <SwiperSlide>
      <img
        src="https://source.unsplash.com/featured/?frontend,developer"
        alt="Frontend Developer"
        className="w-full h-[300px] object-cover"
      />
    </SwiperSlide>
    <SwiperSlide>
      <img
        src="https://source.unsplash.com/featured/?ui,ux"
        alt="UI UX Designer"
        className="w-full h-[300px] object-cover"
      />
    </SwiperSlide>
  </Swiper>
</motion.div>

{/* Backend Section */}
<motion.div
  initial={{ opacity: 0 }}
  whileInView={{ opacity: 1 }}
  transition={{ duration: 0.8 }}
  viewport={{ once: true }}
  className="max-w-4xl mx-auto"
>
  <h2 className="text-2xl font-bold mb-2 text-center text-white"> Backend Team</h2>
  <Swiper
    spaceBetween={30}
    effect={"fade"}
    modules={[EffectFade, Autoplay]}
    autoplay={{ delay: 3000 }}
    loop={true}
    className="rounded-lg overflow-hidden shadow-lg"
  >
    <SwiperSlide>
      <img
        src="https://source.unsplash.com/featured/?backend,developer"
        alt="Backend Developer"
        className="w-full h-[300px] object-cover"
      />
    </SwiperSlide>
    <SwiperSlide>
      <img
        src="https://source.unsplash.com/featured/?api,code"
        alt="API Developer"
        className="w-full h-[300px] object-cover"
      />
    </SwiperSlide>
  </Swiper>
</motion.div>


    </div>
  );
};

export default About;
