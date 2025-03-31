// src/components/About/About.js
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
    <div className="bg-white text-zinc-900 px-6 py-12 max-w-6xl mx-auto space-y-20">
      {/* Hero */}
      <motion.div
        initial={{ opacity: 0, y: 50 }}
        whileInView={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.6 }}
        viewport={{ once: true }}
        className="text-center"
      >
        <h1 className="text-4xl md:text-5xl font-bold">About RoastWear</h1>
        <p className="mt-4 text-lg text-zinc-600 max-w-2xl mx-auto">
          RoastWear is where creativity meets craftsmanship. We help you bring
          your style to life on fabric.
        </p>
      </motion.div>

      {/* Features Section (Cards) */}
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
            className="bg-white shadow-lg rounded-xl p-6 border border-zinc-200 hover:shadow-xl transition duration-300"
          >
            <h3 className="text-xl font-semibold mb-2">{item.title}</h3>
            <p className="text-zinc-500">{item.desc}</p>
          </div>
        ))}
      </motion.div>

      {/* Image Slider Section */}
      <motion.div
        initial={{ opacity: 0 }}
        whileInView={{ opacity: 1 }}
        transition={{ duration: 0.8 }}
        viewport={{ once: true }}
        className="max-w-4xl mx-auto"
      >
        <h2 className="text-2xl font-bold mb-4 text-center">🖼️ Our Journey</h2>
        <Swiper
          spaceBetween={30}
          effect={"fade"}
          modules={[EffectFade, Autoplay]}
          autoplay={{ delay: 3500 }}
          loop={true}
          className="rounded-lg overflow-hidden shadow-lg"
        >
          <SwiperSlide>
            <img
              src="https://source.unsplash.com/featured/?teamwork"
              alt="Team"
              className="w-full h-[300px] object-cover"
            />
          </SwiperSlide>
          <SwiperSlide>
            <img
              src="https://source.unsplash.com/featured/?fashion"
              alt="Design Studio"
              className="w-full h-[300px] object-cover"
            />
          </SwiperSlide>
          <SwiperSlide>
            <img
              src="https://source.unsplash.com/featured/?printing"
              alt="Printing Process"
              className="w-full h-[300px] object-cover"
            />
          </SwiperSlide>
        </Swiper>
      </motion.div>

      {/* CTA Section */}
      <motion.div
        initial={{ opacity: 0, y: 50 }}
        whileInView={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.5 }}
        viewport={{ once: true }}
        className="text-center"
      >
        <h2 className="text-2xl font-semibold mb-2">
          Ready to design your own?
        </h2>
        <p className="text-zinc-600 mb-4">
          Your story deserves to be seen. Let’s print it.
        </p>
        <button className="bg-black text-white px-6 py-2 rounded-lg hover:bg-zinc-800 transition">
          Start Designing
        </button>
      </motion.div>
    </div>
  );
};

export default About;
