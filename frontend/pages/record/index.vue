<template>
  <main class="relative h-screen overflow-hidden font-mono bg-white">
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
    <div class="relative z-20 flex items-center py-10">
      <div
        class="container relative flex flex-col items-center justify-between px-6 py-4 mx-auto"
      >
        <div class="flex flex-col justify-center">
          <h2
            class="max-w-3xl py-2 mx-auto text-5xl font-bold text-gray-800 md:text-3xl"
          >
            WakeUpTime　　{{ "20:00" }}
          </h2>
          <div class="inline-block relative w-72 mt-10">
            <select
              class="block appearance-none w-full bg-white border border-gray-400 hover:border-gray-500 px-4 py-2 pr-8 rounded shadow leading-tight focus:outline-none focus:shadow-outline"
            >
              <option>select Todo</option>
              <option>Option 2</option>
              <option>Option 3</option>
            </select>
            <div
              class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700"
            >
              <svg
                class="fill-current h-4 w-4"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 20 20"
              >
                <path
                  d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z"
                />
              </svg>
            </div>
          </div>
          <div
            class="bg-gray-20 border-8 border-gray-600 p-2 rounded-full h-72 w-72 flex items-center justify-center shadow-lg mt-20"
          >
            <h1 class="text-5xl text-gray-800">00:00:00</h1>
          </div>
          <div class="flex items-center justify-center">
            <div
              v-if="true"
              class="px-4 py-2 m-2 text-gray-800 uppercase bg-transparent border-2 border-gray-800 mt-6 hover:bg-gray-800 hover:text-white text-md"
            >
              START
            </div>
            <div
              v-if="false"
              class="px-4 py-2 text-gray-800 uppercase bg-transparent border-2 border-gray-800 mt-6 hover:bg-gray-800 hover:text-white text-md"
            >
              STOP
            </div>
            <div
              class="px-4 py-2 m-2 text-gray-800 uppercase bg-transparent border-2 border-gray-800 mt-6 hover:bg-gray-800 hover:text-white text-md"
            >
              RESET
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>
