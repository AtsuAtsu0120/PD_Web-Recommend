import { Box} from "@yamada-ui/react"
import { Image } from "@yamada-ui/react"
import { Text } from "@yamada-ui/react"
import { Link } from "@yamada-ui/react"
import { Center } from "@yamada-ui/react"
import { UIProvider } from "@yamada-ui/react"

type BoxProps = {
    imgsrc?: string;
    href: string;
    title: string;
}

export function Box_a({imgsrc, href, title}: BoxProps) { 
  return(
    <Link href={href}>
      <UIProvider>
        <Center>
          <Box w="95%" h="100%" p="md" rounded="8" color="Black" bg="White">
            <Center>
              <Image src={imgsrc} fallback="Frontend/public/error.svg" size="xl" alt="画像"/>
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