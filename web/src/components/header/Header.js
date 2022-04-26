import { PageHeader, Button } from "antd";
import useApp from "../../hooks/useApp";

const Header = () => {
  const { title, func, accessToken } = useApp();

  return (
    <PageHeader
      title={title}
      onBack={() => window.history.back()}
      extra={[
        accessToken && <Button onClick={func.signOut} key="1" type="default">Sign out</Button>
      ]}
    >
    </PageHeader>
  )
}

export default Header;
