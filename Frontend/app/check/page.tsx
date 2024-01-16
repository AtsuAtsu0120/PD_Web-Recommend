"use client"

import { HeroSection } from '@/ui/HeroSection-Explore'
import { RadiusButton } from '@/ui/RadiusButton'
import { Checkbox, Text, Wrap } from "@yamada-ui/react"
import Image from 'next/image'
import { ChangeEventHandler, useState } from 'react'

const inshoData = [
  '好き・嫌い',
  '美しい・汚い',
  'かっこいい・かわいい',
  '明るい・暗い',
  '大人っぽい・子供っぽい',
  '派手・地味',
] 

export default function Check(){
  const [impressionBools, setImpressionBools] = useState([false, false, false, false, false, false])

  const onChangeToggle = (index: number) => {
    const newImpressionBools = [...impressionBools];
    newImpressionBools[index] = !newImpressionBools[index];
    setImpressionBools(newImpressionBools);
  }
  return(
  <div>
    <HeroSection/>
    <div className='flex flex-col gap-12 text-xl pl-20 pt-20'>
    <Text>以下から検索したい印象を選択</Text>
      {inshoData.map((item, index) => (
        // eslint-disable-next-line react/jsx-key
        <div className="flex items-center mb-4">
          <input id="default-checkbox" type="checkbox" value="" className="w-6 h-6 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600" onChange={() => onChangeToggle(index)}/>
          <label htmlFor="default-checkbox" className="ms-2 font-medium text-gray-900 dark:text-gray-300 text-2xl">{item}</label>
        </div>
      ))}
    </div>
    <div className='flex justify-center py-20'>
      <RadiusButton href='/slide' className='bg-gradient-to-tr from-blue-400 to-pink-300'>
        <Image src="/next_grad.svg" alt='ゲーミングNEXT' width={100} height={100}/>
      </RadiusButton>
    </div>
  </div>
  )
}