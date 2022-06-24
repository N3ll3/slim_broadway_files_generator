<?php
declare(strict_types=1);

namespace Macompta\Api\{%BASE_PATH%}\Domain\{%DOMAIN%}\Command;


use Macompta\I18n\I18n;
use Psr\Log\LoggerInterface;

use Macompta\Api\{%BASE_PATH%}\Domain\{%DOMAIN%}\Command\{%NAME%}Command;


class {%NAME%}CommandHandler
{

    public function __construct(
        LoggerInterface $logger,
        DomainRepository $repo
    ) {
        $this->logger = $logger;
        $this->i18n = I18n::load('fr/messages.yaml');
    }

    public function handle({%NAME%}Command $command)
    {
        
        $domain = $this->repo->load($command->aggregateRootId);

        $domain->{%NAME%}(
            $command->aggregateRootId,
            {%PARAMS%}
            );
        $this->repo->save($domain);
    }
}
