package thrift4go.enumm;


import org.apache.thrift.TException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import thrift4go.generated.enumm.ContainerOfEnums;
import thrift4go.generated.enumm.ContainerOfEnumsTestService;


public class EnummServiceDefinition implements ContainerOfEnumsTestService.Iface {
  private static final Logger log = LoggerFactory.getLogger(EnummServiceDefinition.class);

  private final String protocolName;

  public EnummServiceDefinition(final String protocolName) {
    this.protocolName = protocolName;
  }

  @Override
  public ContainerOfEnums echo(final ContainerOfEnums message) throws TException {
    log.info("Echo Service for '{}' received'{}' and will respond with '{}'.",
        new Object[] {protocolName, message, message});

    return message;
  }
}
