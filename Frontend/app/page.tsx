
import { RadiusButton } from '@/ui/RadiusButton'
import { Box, Button } from '@yamada-ui/react'
import Image from 'next/image'

export default function Home() {
  return (
    <div className='flex justify-center items-center min-h-screen text-center'>
      <div>
        <div className='text-3xl py-5'>Aptitude Style Advisor</div>
        <div className='text-7xl py-5 italic'>ASA</div>
        <div className='text-2xl pt-5 pb-2 font-thin'>あなたの望む</div>
        <div className='text-2xl pb-5 font-thin'>サイトが見つかります。</div>
          <RadiusButton href='/check' className='my-5 bg-gradient-to-tr from-blue-400 to-pink-300'>
            <div className='text-2xl bg-clip-text leading-none text-gray-900 bg-gradient-to-tr from-blue-400 to-pink-300 text-transparent'>
              Exprole
            </div>
          </RadiusButton>
      </div>
    </div>
  )
}
