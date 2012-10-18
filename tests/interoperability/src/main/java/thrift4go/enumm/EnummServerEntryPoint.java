package thrift4go.enumm;


import org.apache.thrift.protocol.TBinaryProtocol;
import org.apache.thrift.protocol.TCompactProtocol;
import org.apache.thrift.protocol.TJSONProtocol;
import org.apache.thrift.protocol.TProtocolFactory;
import org.apache.thrift.protocol.TSimpleJSONProtocol;
import org.apache.thrift.server.TServer;
import org.apache.thrift.server.TServer.Args;
import org.apache.thrift.server.TSimpleServer;
import org.apache.thrift.transport.TServerSocket;
import org.apache.thrift.transport.TServerTransport;
import org.apache.thrift.transport.TTransportException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import thrift4go.generated.enumm.ContainerOfEnumsTestService;


public class EnummServerEntryPoint {
  private static final Logger log = LoggerFactory.getLogger(EnummServerEntryPoint.class);

  public static void start(String protocol, int port) throws TTransportException {
    log.info("Preparing to start echo service.");

    final ContainerOfEnumsTestService.Processor<EnummServiceDefinition> processor =
        new ContainerOfEnumsTestService.Processor<EnummServiceDefinition>(
            new EnummServiceDefinition(protocol));
    final TServerTransport transport = new TServerSocket(port);

    final Args serviceArguments = new Args(transport);
    serviceArguments.processor(processor);
    serviceArguments.protocolFactory(Enum.valueOf(Protocol.class, protocol).getFactory());

    final TServer server = new TSimpleServer(serviceArguments);

    log.info("Provisioned everything; now serving {} requests on {}...", protocol,port);

    try {
      server.serve();
    } finally {
      log.info("Closing down everything.");

      server.stop();
    }
  }

  private static enum Protocol {
    JSON(new TJSONProtocol.Factory()),
    SIMPLE_JSON(new TSimpleJSONProtocol.Factory()),
    BINARY(new TBinaryProtocol.Factory()),
    COMPACT(new TCompactProtocol.Factory());

    private final TProtocolFactory factory;

    Protocol(final TProtocolFactory factory) {
      this.factory = factory;
    }

    public TProtocolFactory getFactory() {
      return this.factory;
    }
  }
}
