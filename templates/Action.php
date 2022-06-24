<?php

declare(strict_types=1);
namespace Macompta\Api\{%BASE_PATH%}\Application\Actions\{%DOMAIN%};

use Psr\Log\LoggerInterface;
use \Slim\Http\Response as Response;
use Macompta\Slim\Bridge\Application\Actions\Action;
use Macompta\Broadway\Bridge\CommandHandling\CommandBus;

use Macompta\Api\{%BASE_PATH%}\Domain\{%DOMAIN%}\Command\{%NAME%}Command;


class {%NAME%}Action extends Action
{
    public function __construct(
        LoggerInterface $logger,
        CommandBus $cb
    ) {
        parent::__construct($logger);
        $this->cb = $cb;
    }

    protected function action(): Response
    {
        try {
            $data = $this->request->getParsedBody();
            
            ${%DOMAIN%}Id = new {%DOMAIN%}Id();

            $command = new {%NAME%}Command(${%DOMAIN%}Id, $data);
            $this->cb->dispatch($command);

            $this->logger->info(json_encode([
                'compte' => '',
                'compte_uuid' =>'',
                'canal'   => '',
                'status'  => 'ok',
                'msg'     => '',
                'payload' => $data,
                'referer' => '',
                'route'   => '',
                'http_code' => ''
            ]));
            return $this->response->withJson('OK', 200);
        } catch (\PDOException $error) {
            $this->logger->error(json_encode([
                'compte' => '',
                'compte_uuid' => '',
                'canal'   => '',
                'status'  => 'ko',
                'msg'     => $error->getMessage()." - ligne : ".$error->getLine()."- file :".$error->getFile(),
                'payload' => ["sql_code" => $error->getCode() ],
                'referer' => '',
                'route'   => '',
                'http_code' => '400'
            ]));
            return $this->responseException($error);
        } catch (\Exception $e) {
            $this->logger->error(json_encode([
                'compte' => '',
                'compte_uuid' => '',
                'canal'   => '',
                'status'  => 'ko',
                'msg'     => $e->getMessage(),
                'payload' => $data,
                'referer' => '',
                'route'   => ''
            ]));
            return $this->responseException($e);
        }
    }

}
