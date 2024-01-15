import { Box} from "@yamada-ui/react"
import { Text } from "@yamada-ui/react"
import { Link } from "@yamada-ui/react"
import { Center } from "@yamada-ui/react"
import { UIProvider } from "@yamada-ui/react"
import Image from "next/image"

type BoxProps = {
    imgsrc: string;
    href: string;
    title: string;
}

export function Box_a({imgsrc, href, title}: BoxProps) { 
  return(
    <Link href={href} className="w-11/12 py-10">
      <UIProvider>
        <Center>
          <Box w="95%" h="100%" p="md" rounded="8" color="Black" bg="White">
            <Center>
              <Image src={imgsrc} alt="ホームページのスクリーンショット" width={700} height={700} className="shadow-xl"/>
            </Center>
            <Box m="5%">
              <Text fontWeight="extrabold" fontSize="100xl" isTruncated>
                {title}
              </Text>
              <Link href={href} isExternal>
                {href}
              </Link>
              </Box>
          </Box>
        </Center>
      </UIProvider>
    </Link>
    );
    }