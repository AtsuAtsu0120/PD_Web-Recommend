import { Box} from "@yamada-ui/react"
import { Image } from "@yamada-ui/react"
import { Text } from "@yamada-ui/react"
import { Link } from "@yamada-ui/react"
import { Center } from "@yamada-ui/react"
import { UIProvider } from "@yamada-ui/react"

export function Box_a() { 
    return(
    <UIProvider>
     <Center>
     <Box w="95%" h="100%" p="md" rounded="8" color="Black" bg="White">
      <Center>
        <Image src="Frontend/public/error.svg" fallback="Frontend/public/error.svg" size="xl" alt="画像"/>
      </Center>
      <Box m="5%">
      <Text fontWeight="extrabold" fontSize="100xl" isTruncated>
        株式会社 CirKit
      </Text>
      <Link
      href="https://cirkit.jp/" isExternal>
        https://cirkit.jp/
      </Link>
      </Box>
     </Box>
     </Center>
     </UIProvider>
     );
  }