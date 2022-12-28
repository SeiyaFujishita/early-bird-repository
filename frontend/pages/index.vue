<template>
  <main
    class="relative h-screen overflow-hidden font-mono bg-white dark:bg-gray-800"
  >
    <div class="absolute hidden md:block -bottom-32 -left-32 w-96 h-96">
      <div
        class="absolute z-20 text-xl text-extrabold right-12 text-start top-1/4"
      >
        <span class="text-7xl"> 🌱 </span>
        <p>おはようございます</p>
      </div>
      <svg
        viewBox="0 0 200 200"
        class="absolute w-full h-full"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          fill="#FFDBB9"
          d="M44.7,-76.4C58.8,-69.2,71.8,-59.1,79.6,-45.8C87.4,-32.6,90,-16.3,88.5,-0.9C87,14.6,81.4,29.2,74.1,43.2C66.7,57.2,57.6,70.6,45,78.1C32.4,85.6,16.2,87.1,0.7,85.9C-14.8,84.7,-29.6,80.9,-43.9,74.4C-58.3,67.9,-72,58.7,-79.8,45.9C-87.7,33,-89.5,16.5,-88.9,0.3C-88.4,-15.9,-85.4,-31.7,-78.1,-45.4C-70.8,-59.1,-59.1,-70.6,-45.3,-77.9C-31.6,-85.3,-15.8,-88.5,-0.3,-88.1C15.3,-87.6,30.5,-83.5,44.7,-76.4Z"
          transform="translate(100 100)"
        ></path>
      </svg>
    </div>
    <div class="relative z-20 flex items-center py-60">
      <div
        class="container relative flex flex-col items-center justify-between px-6 py-4 mx-auto"
      >
        <div class="flex flex-col">
          <h2
            class="max-w-3xl py-2 mx-auto text-5xl font-bold text-center text-gray-800 md:text-3xl dark:text-white"
          >
            ~ 早起きは勤勉ということで人から信頼が得られる ~
          </h2>
          <div class="flex items-center justify-center mt-4">
            <NuxtLink
              to="http://localhost:3000/record"
              class="px-4 py-2 my-2 text-gray-800 uppercase bg-transparent border-2 border-gray-800 md:mt-16 dark:text-gray-800 dark:bg-white hover:dark:bg-gray-100 dark:text-white hover:bg-gray-800 hover:text-white text-md"
            >
              Wake Up !
            </NuxtLink>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>
