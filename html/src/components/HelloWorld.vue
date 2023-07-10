<template>
  <div class="root">
    <!--导航条和头部内容-->
    <div class="box">
      <div class="content">
        <!--菱形斜面-->
        <svg class="diamond" viewBox="0 0 100 100" preserveAspectRatio="none">
          <polygon points="50,0 100,0 50,100 0,100"/>
        </svg>
        <div>
          <!--大屏幕横向导航-->
          <div class="flat">
            <nav class="nav">
              <div class="parts">
                <div class="logo">
                  <a href="#"><img src="https://tailwindui.com/img/logos/workflow-mark-indigo-600.svg" alt=""/></a>
                  <div class="toggle">
                    <button type="button">
                      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                              d="M4 6h16M4 12h16M4 18h16"/>
                      </svg>
                    </button>
                  </div>
                </div>
              </div>
              <div class="menu">
                <a href="#">Product</a>
                <a href="#">Features</a>
                <a href="#">Marketplace</a>
                <a href="#">Company</a>
                <a href="#" class="login">Log in</a>
              </div>
            </nav>
          </div>
          <!--小屏幕弹窗导航-->
          <!--          <div class="modal">-->
          <!--            <div class="rounded-lg shadow-md bg-white ring-1 ring-black ring-opacity-5 overflow-hidden">-->
          <!--              <div class="px-5 pt-4 flex items-center justify-between">-->
          <!--                <div>-->
          <!--                  <img class="h-8 w-auto" src="https://tailwindui.com/img/logos/workflow-mark-indigo-600.svg" alt=""/>-->
          <!--                </div>-->
          <!--                <div class="-mr-2">-->
          <!--                  <button type="button"-->
          <!--                          class="bg-white rounded-md p-2 inline-flex items-center justify-center text-gray-400 hover:text-gray-500 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-indigo-500">-->
          <!--                    <svg class="h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"-->
          <!--                         stroke="currentColor" aria-hidden="true">-->
          <!--                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>-->
          <!--                    </svg>-->
          <!--                  </button>-->
          <!--                </div>-->
          <!--              </div>-->
          <!--              <div class="px-2 pt-2 pb-3 space-y-1">-->
          <!--                <a href="#"-->
          <!--                   class="block px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-50">Product</a>-->
          <!--                <a href="#"-->
          <!--                   class="block px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-50">Features</a>-->
          <!--                <a href="#"-->
          <!--                   class="block px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-50">Marketplace</a>-->
          <!--                <a href="#"-->
          <!--                   class="block px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-50">Company</a>-->
          <!--              </div>-->
          <!--              <a href="#"-->
          <!--                 class="block w-full px-5 py-3 text-center font-medium text-indigo-600 bg-gray-50 hover:bg-gray-100">Log-->
          <!--                in</a>-->
          <!--            </div>-->
          <!--          </div>-->
        </div>
        <main class="main">
          <div class="sm:text-center lg:text-left">
            <h1 class="text-4xl tracking-tight font-extrabold text-gray-900 sm:text-5xl md:text-6xl">
              <span class="block xl:inline">Data to enrich your</span>
              <span class="block text-indigo-600 xl:inline">online business</span>
            </h1>
            <p class="mt-3 text-base text-gray-500 sm:mt-5 sm:text-lg sm:max-w-xl sm:mx-auto md:mt-5 md:text-xl lg:mx-0">
              Anim aute id magna aliqua ad ad non deserunt sunt. Qui irure qui lorem cupidatat commodo. Elit sunt amet
              fugiat veniam occaecat fugiat aliqua.
            </p>
            <div class="mt-5 sm:mt-8 sm:flex sm:justify-center lg:justify-start">
              <div class="rounded-md shadow">
                <a href="#"
                   class="w-full flex items-center justify-center px-8 py-3 border border-transparent text-base font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 md:py-4 md:text-lg md:px-10">
                  Get started
                </a>
              </div>
              <div class="mt-3 sm:mt-0 sm:ml-3">
                <a href="#"
                   class="w-full flex items-center justify-center px-8 py-3 border border-transparent text-base font-medium rounded-md text-indigo-700 bg-indigo-100 hover:bg-indigo-200 md:py-4 md:text-lg md:px-10">
                  Live demo
                </a>
              </div>
            </div>
          </div>
        </main>
      </div>
    </div>
    <!--装饰图-->
    <div class="decoration">
      <img
          src="https://images.unsplash.com/photo-1551434678-e076c223a692?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=2850&q=80"
          alt=""/>
    </div>
  </div>
</template>
<style scoped lang="scss">
.root {
  @apply relative bg-white overflow-hidden;
  .box {
    @apply max-w-7xl mx-auto;
    .content {
      @apply relative z-10 pb-8 bg-white sm:pb-16 md:pb-20 lg:max-w-2xl lg:w-full lg:pb-28 xl:pb-32;
      .diamond {
        @apply hidden lg:block absolute right-0 inset-y-0 h-full w-48 text-white transform translate-x-1/2 fill-current;
      }

      .flat {
        @apply relative pt-6 px-4 sm:px-6 lg:px-8;
        .nav {
          @apply relative flex items-center justify-between sm:h-10 lg:justify-start;
          .parts {
            @apply flex items-center flex-grow flex-shrink-0 lg:flex-grow-0;
            .logo {
              @apply flex items-center justify-between w-full md:w-auto;
              img {
                @apply h-8 w-auto sm:h-10;
              }

              .toggle {
                @apply -mr-2 flex items-center md:hidden;
                button {
                  @apply bg-white rounded-md p-2 inline-flex items-center justify-center text-gray-400 hover:text-gray-500 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-indigo-500;
                }

                svg {
                  @apply h-6 w-6;
                }
              }
            }
          }

          .menu {
            @apply hidden md:block md:ml-10 md:pr-4 md:space-x-8;
            a {
              @apply font-medium text-gray-500 hover:text-gray-900;
            }

            .login {
              @apply font-medium text-indigo-600 hover:text-indigo-500;
            }
          }
        }
      }

      .modal {
        @apply absolute z-10 top-0 inset-x-0 p-2 transition transform origin-top-right md:hidden;
      }

      .main {
        @apply mt-10 mx-auto max-w-7xl px-4 sm:mt-12 sm:px-6 md:mt-16 lg:mt-20 lg:px-8 xl:mt-28;
      }
    }
  }

  .decoration {
    @apply lg:absolute lg:inset-y-0 lg:right-0 lg:w-1/2;
    img {
      @apply h-56 w-full object-cover sm:h-72 md:h-96 lg:w-full lg:h-full;
    }
  }
}
</style>
<script lang="ts" setup>

</script>